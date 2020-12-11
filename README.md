# Go-gituser

This programs automates the git config command. <br>
So you can change between git users accouts easily.

#### Modes

There is currently 3 modes in this script:

- "work" : for work related git account/
- "school" : for school related git account.
- "personal" : for personal relared git account.

### Configuration

To add your respective accounts, you need to fill out the `data/config.json` file.

```
{
  "personalUsername": "enterYourUsernameHere",
  "personalEmail": "enterYourEmailHere",
  "schoolUsername": "enterYourUsernameHere",
  "schoolEmail": "enterYourEmailHere",
  "workUsername": "enterYourUsernameHere",
  "workEmail": "enterYourEmailHere"
}

```

## Usage

<i>Attention: </i> Make sure you've entered your information in `config.json` before compile program

Compile by running `go build -o gituser`

Call executable file with mode

```
./gituser <mode>
```

Examples:

```
./gituser work
```

```
./gituser school
```

```
./gituser personal
```

### Help

There is also a flag `-help` that will print some information about the program.

`./gituser -help`
