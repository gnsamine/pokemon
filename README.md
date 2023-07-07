# pokemon
Ghost-type Pokémon Sprite Downloader.

This Go program allows you to fetch the names and front-default sprites of Ghost-type Pokémon from the PokeAPI and save the front-default sprites as PNG images.

# Getting Started

To get started with the Pokémon Sprite Downloader, follow these steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/pokemon-sprite-downloader.git
Navigate to the project directory:
    cd pokemon-sprite-downloader

Install the required dependencies:
    go get -u github.com/gnsamine/pokemon/poke

Run the program:
    go run main.go


#Usage
The program utilizes the PokeAPI to retrieve the list of Ghost-type Pokémon and their respective sprites. It performs the following steps:

Fetches the list of Ghost-type Pokémon from the PokeAPI.
Iterates through each Pokémon in the list.
Retrieves the URL of the front default sprite for each Pokémon.
Downloads the sprite image and saves it as a PNG file in the images/ directory.

#Contributing
Contributions to the Pokémon Sprite Downloader project are welcome! If you would like to contribute, please follow these steps:

Fork the repository.
Create a new branch for your feature or bug fix.
Make your changes.
Commit your changes and push them to your forked repository.
Submit a pull request, describing your changes in detail.
Please ensure that your contributions adhere to the project's coding conventions and follow good software development practices.

#License
This project is licensed under the MIT License.
Make sure to replace "your-username" in the repository URL (`https://github.com/your-username/pokemon-sprite-downloader.git`) with your actual GitHub username or the appropriate repository URL.
Feel free to customize the README file further with additional sections, examples, or details based on your project's requirements and specific features.
