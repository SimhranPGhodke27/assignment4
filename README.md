# MyDB Project

This README provides instructions for setting up and running the MyDB project, which utilizes Go, GORM, and PostgreSQL. Follow these steps to configure and launch the application on your local machine.

## 1. Set Up Your Go Environment

- **Ensure Go is Installed**: Make sure you have Go installed and set up on your machine. If Go is not installed, follow the [Go installation instructions](https://golang.org/doc/install) to get it up and running.

## 2. Install GORM and Database Driver

- **Install GORM and PostgreSQL Driver**: You need GORM and a database driver for your database.  Install GORM and the sqlite driver using the following commands:

  ```bash
  go get -u gorm.io/gorm
  go get -u gorm.io/driver/sqlite
  ```

## 3. Set Up the Project Directory

1. **Create a Project Folder**
   - Create a folder named `mydb` in your local directory.

2. **Download Repository Files**
   - Download all the files from this repository into the `mydb` folder you created.

3. **Open Command Prompt and Navigate to Project Directory**
   - Open Command Prompt on your computer.
   - Change the directory to the `mydb` project folder using the following command:

     ```bash
     cd your\projectdirectory\mydb
     ```

## 4. Run the Application

- **Start the Application**: Run the following command to start the application:

  ```bash
  go run main.go
  ```

- **Application Output**: You should see messages indicating that a user was created, read, updated, and deleted.
