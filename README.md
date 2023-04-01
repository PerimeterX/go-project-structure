# Go Project Structure

<img align="right" width="320" alt="go-project-structure-logo" src="https://raw.githubusercontent.com/PerimeterX/go-project-structure/assets/logo.svg">

[![Licence](https://img.shields.io/github/license/perimeterx/go-project-structure)](LICENSE)
[![Latest Release](https://img.shields.io/github/v/release/perimeterx/go-project-structure)](https://github.com/perimeterx/go-project-structure/releases)
[![Issues](https://img.shields.io/github/issues/perimeterx/go-project-structure?logo=github)](https://github.com/perimeterx/go-project-structure/issues)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)](CODE_OF_CONDUCT.md)

This is a template repository for Go projects.

We released a [blog post](https://www.humansecurity.com/tech-engineering-blog/finding-the-best-go-project-structure-part-1) revealing the story behind this structure. 
Our internal journey at HUMAN Security to find the best project structure, the decision we’ve taken,
and the conclusions we’ve drawn along the way.

To view a full example, navigate to the [example branch](https://github.com/PerimeterX/go-project-structure/tree/example).

Each of the directories in this repository contains a README file describing its purpose.

## Contents

- [Usage](#usage)
- [Project Structure Design](#project-structure-design)
- [Contact and Contribute](#contact-and-contribute)

## Usage

To create a new project, fork this repository or [use it as a template repository](https://github.com/PerimeterX/go-project-structure/generate).

## Project Structure Design

The project structure is designed with [independent packages](https://medium.com/@avivcarmis/ok-lets-go-three-approaches-to-structuring-go-code-42e2370c3ae5#92df) in mind,
according to [hexagonal architecture](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)), and targeted to provide the following benefits:
- Structural coherence
- Separation of business logic and infrastructure
- Reusability
- Dependency declaration

Fore more info, read the [full article](https://www.humansecurity.com/tech-engineering-blog/finding-the-best-go-project-structure-part-1) behind this structure.

## Contact and Contribute

Reporting issues and requesting features may be done in our [GitHub issues page](https://github.com/PerimeterX/go-project-structure/issues).
For any further questions or comments you can reach us out at [open-source@humansecurity.com](mailto:open-source@humansecurity.com).

Any type of contribution is warmly welcome and appreciated ❤️
Please read our [contribution](CONTRIBUTING.md) guide for more info.
