# MH-API (Monster Hunter API)

MH-API is an open source project that provides strategy information and other information related to the Monster Hunter series. This project is developing an API to allow players of the Monster Hunter series to access game strategy information, etc., and to easily create secondary works and tools, etc.

## Introduction

This README.md describes the guidelines and usage of the MH-API project. You are welcome to participate in the project according to the following guidelines.

The Code of Conduct for this project can be found [here](. /CODE_OF_CONDUCT.md)

## Getting Started

To join the MH-API project, please follow the steps below.

### Preparation

- Review the Contribute Guide. [Click here for the Contribute Guide](./CONTRIBUTING.md)

### Set up the environment

1. go to the repository directory

    ```bash
        cd MH-API
    ```

2. Open the directory in an editor
3. Create a new branch.

    ```bash
        git checkout -b "[new branch]"
    ```

4. Check that it works.

   ```bash
        # start up docker
        make up

        # Make sure the response is {"message": "ok"}.
        curl http://localhost:8080/v1/system/healthcheck
   ````

5. Run the test

    ```bash
        # All tests should return "ok".
        make test
    ```

## Communicate with the community

In this community, you can find the CODE_OF_CONDUCT.md file [here](/CODE_OF_CONDUCT.md)

To participate in the MH-API project, the following communication channels are available

- Slack channel: join [slack.mhapi.org](https://slack.mhapi.org) to interact with other contributors and members.

- Issue Tracker: Use the [MH-API Issue Tracker](https://github.com/mhapi/issues) to report bugs and suggest new features.

- Mailing Lists: Join [mhapi-dev@groups.com](mailto:mhapi-dev@groups.com) to receive discussions and important announcements via email.

## License

The MH-API project is released under the [MIT License](https://opensource.org/licenses/MIT). Please see the LICENSE file in the project for detailed license information.

## Contribution Guidelines

Please refer to [CONTRIBUTING.md](./CONTRIBUTING.md) for guidelines on contributing to the MH-API project. Please read the guidelines before contributing code or documentation to the project.

## SUPPORT

If you need support regarding the MH-API project, please contact [support@mhapi.org](mailto:support@mhapi.org).

## Acknowledgements

This project is made possible by contributors from the open source community and MH-API users. Many people deserve our thanks.

For more information about the project and updates, please visit the [official MH-API website](https://mhapi.org).

This project uses the Monster Hunter Series™, a trademark and registered trademark of Capcom Co. Ltd . The Monster Hunter Series™ is the intellectual property of Capcom Co. Ltd . We hereby express our gratitude to Capcom.

This project is unofficial and has nothing to do with Capcom Co Ltd .

**Happy coding!

### References

<https://opensource.guide/ja/starting-a-project/>

This project has started from 2023/5/21
