# Wriked

A simple cli to upload hours to a Wrike task from an Excel (xlsx) spreadsheet

[![Build Status](https://travis-ci.org/andymotta/wriked.svg?branch=master)](https://travis-ci.org/andymotta/wriked)

## Pre-Installation
Generate and store your Wrike API token in a safe place
https://developers.wrike.com/documentation/oauth2#skipoauth
('Permanent access token' section)

## Installation

[Download the latest binary](https://github.com/andymotta/wriked/releases) for your architecture and be sure to make it executable

For example:  `~/Downloads> chmod +x wriked-macOS-amd64`


## Usage

```bash
./wriked-macOS-amd64 -a <your_api_token> -f ~/Downloads/TimeSheet-August\ 2018.xlsx
# Filenames with spaces need to be escaped

  -a string
    	Get Dangerous with your Wrike API Token
  -f string
    	Excel file containing hours to upload to Wrike
```

## Support

Please [open an issue](https://github.com/andymotta/wriked/issues/new) for support.

## Contributing

Please contribute using [Github Flow](https://guides.github.com/introduction/flow/). Create a branch, add commits, and [open a pull request](https://github.com/andymotta/wriked/compare/).
