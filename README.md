<p align="center">
  <img src="https://github.com/user-attachments/assets/20c1bf00-2e5a-4280-b891-b9243c14fb89" width="300" />
</p/>

<h1 align="center">ArmoraCrypt</h1>

<p>
  A lightweight, fast CLI tool for encrypting and storing files securely on the cloud.<br>Written in Go with <strong><a href="https://cobra.dev">cobra</a> </strong>
</p>

## Features

-  **AES-256 Encryption** - Military-grade encryption for all file types
-  **Folder Support** - Automatically zip and encrypt entire directories
-  **Dropbox Integration** - Seamlessly upload and download encrypted files
-  **Local Security** - Keys never leave your machine
-  **Easy to Use** - Simple CLI commands for all operations

## Quick Start

### Prerequisites
- **Go 1.16** or higher ([Download](https://golang.org/dl/))

### Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/ArmoraCrypt.git
cd ArmoraCrypt

# Install the binary
go install

# Verify installation
armoracrypt
```

That's it! You're ready to encrypt.

## Demo


https://github.com/user-attachments/assets/5d0e5691-54da-46a3-933c-85bffcdac5e1


##  Usage Guide

### Basic Commands

#### 🔐 Encrypt a Single File
```bash
armoracrypt encrypt --fp "path/to/your/document.pdf"
```
Creates `document.pdf.crypt`

#### 📦 Encrypt a Folder
```bash
armoracrypt encrypt --d "path/to/your/folder"
```
Automatically zips and encrypts the entire folder

#### 🔓 Decrypt Files or Folders
```bash
armoracrypt decrypt --fp "path/to/encrypted/file.crypt"
```
Works with both files and encrypted folder archives

#### ☁️ Upload to Dropbox
```bash
armoracrypt upload --fp "path/to/file/folder"
```
Automatically encrypts your file/folder and uploads it to cloud
#### ⬇️ Download from Dropbox
```bash
armoracrypt download
```
Retrieve and decrypt files from your Dropbox

## Configuration

### Setting Up Dropbox Integration

1. **Create a Dropbox App**
   - Go to [Dropbox Developer](https://www.dropbox.com/developers/apps)
   - Create a new app
   - Generate an access token
   - Enter the token when prompted
