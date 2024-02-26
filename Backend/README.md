# Running the Technical Challenge

This repository contains a simple Go program that scans files in different configuration (i.e, All, CrossSiteScripting, SensitiveDataExposure, or SQLInjection)

## Prerequisites

Before running the program, make sure you have Go installed on your machine. You can download and install it from the [official Go website](https://golang.org/).

## How to Run

1. **Clone the Repository:**

    ```bash
    git clone https://github.com/LuisFRosa/CheckmarxGo.git
    ```

2. **Navigate to the Project Directory:**

    ```bash
    cd .\Program\
    ```

3. **Build the Go Program:**

    ```bash
    go build .
    ```

4. **Run the Executable:**

    ```bash
    ./Program main.go ./code All
    ```

## Usage

Once the program is running, it will prompt you to enter the program main. After entering program main (i.e, main.go), then enter the path to the source, it will also run with a scan configuration parameter.

Example of outcome:

```plaintext
[Cross site scripting] in file "checmark.html" on line 25
[Sensitive data exposure.] in file "checmark.html" on line 27
[Sensitive data exposure.] in file "checmark.html" on line 34
[Sensitive data exposure.] in file "checmark.html" on line 41
[SQL injection] in file "checmark2.txt" on line 1
[SQL injection] in file "checmark2.txt" on line 8
[Sensitive data exposure.] in file "checmark2.txt" on line 10