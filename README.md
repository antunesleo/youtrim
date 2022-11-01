docker-compose run youtrim trim --url https://www.youtube.com/watch\?v\=ljeCPM38d8U --start 2 --end 4

go test  ./... -run Unit

go test  ./... -run Integration


<div align="center">
  <h3 align="center">Youtrim</h3>

  <p align="center">
    Download and trim youtube videos!
    .
    <a href="https://github.com/antunesleo/youtrim/issues">Report Bug</a>
    ·
    <a href="https://github.com/antunesleo/youtrim/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
        <li><a href="#installation">Tests</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

Sometimes we need a easy way to trim/crop a youtube video. This tool just do that!

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

* [![Go][Golang]][https://go.dev/]
* [![Moviego][Moviego]][https://github.com/mowshon/moviego]
* [![Youtube][Youtube]][https://github.com/kkdai/youtube]

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

### Prerequisites

* Docker

### Installation

_Below is an example of how you can instruct your audience on installing and setting up your app. This template doesn't rely on any external dependencies or services._

1. Clone the repo
   ```sh
   git clone https://github.com/antunesleo/youtrim
   ```
2. Build docker image
   ```sh
   docker-compose build
   ```
3. Have fun!
   ```sh
   docker-compose run youtrim trim --url https://www.youtube.com/watch\?v\=ljeCPM38d8U --start
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>


### Tests
1. Unit
   ```sh
   docker-compose run --entrypoint='go test  ./... -run Unit' youtrim
   ```
2. Integration
   ```sh
   docker-compose run --entrypoint='go test  ./... -run Integration' youtrim
   ```

<!-- USAGE EXAMPLES -->
## Usage

   ```sh
   docker-compose run youtrim trim --url url --start startSecond --end endSecond
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
## License

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Leo Antunes - antunesleo4@gmail.com

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

<p align="right">(<a href="#readme-top">back to top</a>)</p>
