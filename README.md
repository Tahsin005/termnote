# TermNote - A note taking application, directly in your terminal!

#### TermNote is a simple, elegant, and terminal-based note-taking application built with Bubble Tea and Lipgloss
#### It lets you create, view, edit, and manage your markdown notes directly from the terminal — with smooth keyboard controls and a modern TUI look.

## Features

- Create new notes directly in your terminal
- View a list of existing notes
- Edit and save notes instantly
- Notes stored locally in ~/.termnote
- Beautiful TUI powered by Charm libraries
- Keyboard shortcuts for seamless workflow


## Installation

You can either clone and build TermNote yourself or download the prebuilt executable.

### Option 1 — Build It Yourself

```
git clone https://github.com/Tahsin005/termnote.git
cd termnote
go build -o termnote
```

### Option 2 — Download Prebuilt Executable

If you don’t want to build it yourself, simply download the precompiled executable from Google Drive:

[Download](https://drive.google.com/file/d/10-8ZIdWJzeBVmdUGKw-oFLAckB2ifJ5P/view?usp=drive_link)

- Make sure to give the `termnote` executable, execute permission

### 🐧 Linux: Make Your Go Application Executable

##### 1. Build the Go Application

```
go build -o termnote
```

This creates an executable file named `termnote` in your current directory.

##### 2. Move the Executable to a Directory in Your PATH

```
sudo mv termnote /usr/local/bin/
```

- If you don’t have `sudo` privileges, use:

```
mv termnote ~/.local/bin/
```

##### 3. Ensure the Directory is in Your PATH

If you used `~/.local/bin`, add it to your shell config (e.g., `~/.bashrc` or `~/.zshrc`):

```
export PATH="$HOME/.local/bin:$PATH"
```

Then reload your shell

```
source ~/.bashrc   # or source ~/.zshrc
```

#### 4. Test the Command

```
termnote
```

### 🪟 Windows: Make Your Go Application Executable

#### 1. Build the Go Application

Open PowerShell or Command Prompt in your project folder and run:

```
go build -o termnote.exe
```
This creates an executable file named `termnote.exe`.

##### 2. Move the Executable to a Directory in Your PATH

Move `termnote.exe` to a folder that’s already in your system PATH, for example:

```
C:\Users\<YourName>\AppData\Local\Microsoft\WindowsApps
```

or create your own folder (e.g., `C:\GoApps`) and add it to the PATH.

#### 3. (Optional) Add Folder to PATH

- Press Win + R, type sysdm.cpl, and press Enter.
- Go to Advanced → Environment Variables.
- Under “User variables,” find and edit Path, then add your folder path.

#### 4. Test the Command

```
termnote
```

### 🍏 macOS: Make Your Go Application Executable

##### 1. Build the Go Application

```
go build -o termnote
```

This creates an executable file named `termnote` in your current directory.

##### 2. Move the Executable to a Directory in Your PATH

```
sudo mv termnote /usr/local/bin/
```

- If you prefer not to use `sudo`, use:

```
mv termnote ~/bin/
```

##### 3. Ensure the Directory is in Your PATH

Add this line to your shell config (`~/.zshrc` or `~/.bash_profile`):

```
export PATH="$HOME/bin:/usr/local/bin:$PATH"
```

Then reload your shell

```
source ~/.zshrc
```

#### 4. Test the Command

```
termnote
```

### Exampls

<img src="https://github.com/Tahsin005/termnote/blob/main/assets/ss-1.png">
<img src="https://github.com/Tahsin005/termnote/blob/main/assets/ss-2.png">
<img src="https://github.com/Tahsin005/termnote/blob/main/assets/ss-3.png">
<img src="https://github.com/Tahsin005/termnote/blob/main/assets/ss-4.png">
