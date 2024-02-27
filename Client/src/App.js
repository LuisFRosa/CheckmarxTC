import React, { useState } from 'react';
import axios from 'axios';
import './App.css';

function App() {
  const [selectedFiles, setSelectedFiles] = useState([]);
  const [selectedScanConfig, setSelectedScanConfig] = useState('All');
  const [errors, setErrors] = useState([]);

  const handleFileChange = (event) => {
    setSelectedFiles(event.target.files);
  };

  const handleScanConfigChange = (event) => {
    setSelectedScanConfig(event.target.value);
  };

  const handleSubmit = async () => {
    const formData = new FormData();
    for (let i = 0; i < selectedFiles.length; i++) {
      formData.append('files', selectedFiles[i]);
    }

    try {
      const response = await axios.post('http://localhost:8080/scan', formData, {
        params: {
          scanConfig: selectedScanConfig
        },
        headers: {
          'Content-Type': 'multipart/form-data'
        }
      });

      setErrors(response.data.errors);
    } catch (error) {
      setErrors(['Error scanning files: '+ error]);
    }
  };

  return (
    <div className="App">
      <h1>File Error Scanner</h1>
      <select value={selectedScanConfig} onChange={handleScanConfigChange}>
        <option value="All">All</option>
        <option value="CrossSiteScripting">Cross Site Scripting</option>
        <option value="SensitiveDataExposure">Sensitive Data Exposure</option>
        <option value="SQLInjection">SQL Injection</option>
      </select>
      <input type="file" multiple onChange={handleFileChange} />
      <button onClick={handleSubmit} disabled={selectedFiles.length===0}>Scan Files</button>
      <div className="errors">
        {selectedFiles.length > 0 && errors.length > 0 && (
          <div>
            <h2>Errors:</h2>
            <ul>
              {errors.map((error, index) => (
                <li key={index}>{error}</li>
              ))}
            </ul>
          </div>
        )}
      </div>
    </div>
  );
}

export default App;
