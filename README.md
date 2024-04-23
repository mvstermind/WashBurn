# WashBurn - Discord Midi Generating Bot 

This Discord bot is designed to assist music producers who may struggle with determining the key or scale they want to work with. Bot can generate MIDI chord progression with BPM for easier start,
or for just randomly generated key and BPM.

## Getting Started

To use this bot, follow these steps:

1. **Clone the Repository:** Begin by cloning this repository to your local machine using Git:

    ```
    git clone https://github.com/mvstermind/WashBurn.git
    ```

2. **Configure the Bot Token:**
    - Navigate to the `handler.go` file and locate the `token` variable.
    - Uncomment the `token` variable and add your own Discord bot token. You can obtain a bot token by creating a new bot application on the Discord Developer Portal.

3. **Invite the Bot to Your Discord Server:**
    - Once your bot application is created and the token is configured, invite the bot to your Discord server using the OAuth2 URL provided in the Discord Developer Portal.

4. **Build and Run the Bot:**
    - Build the bot using your preferred Go build tool. For example:
    
    ```
    go build -o WashBurn .
    ```

    - Run the bot executable:

    ```
    ./WashBurn
    ```

## Usage

This bot supports the following commands:

- **-generate** or **-gen**: Generates both random BPM and MIDI chord progression.
- **-bpm** or **-b**: Generates BPM.
- **-s** or **-scale**: Generates scale.
- **-sbpm**: Generates Bpm and scale.

## Contributing

If you have suggestions for improvements or new features, feel free to open an issue or submit a pull request. Your contributions are welcome and appreciated!

## License

This project is licensed under the [MIT License](LICENSE), which means you are free to modify and distribute the code as long as you include the original license and copyright notice.


