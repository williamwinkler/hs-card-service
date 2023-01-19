# Hearthstone Card Service
Serves cards from Blizzards card game [Hearthstone](https://hearthstone.blizzard.com/en-us).

The very purpose of this program was to personally get well acquainted with Swagger codegen, Clean Architecture and MongoDb. <br>
The program gets its data from [Blizzards Hearthstone API](https://develop.battle.net/documentation/hearthstone/game-data-apis) and stores it in a mongoDB database.

## API
The API is defined in `./api/swagger.yml` which is used to generate the server located in `./codegen.` <br>
The API can be seen [here](https://app.swaggerhub.com/apis-docs/WILLIAMBWINKLER/hearthstone-card-service/1.0.0#/) on SwaggerHub.

## Architecture
This program is designed using the principles of Clean Architecture. It emphasizes separation of concerns and adheres to SOLID principles. The application is divided into distinct layers, with each layer having a specific responsibility, allowing for greater flexibility and maintainability of the codebase. The program is designed to be independent of any particular framework or technology, making it a suitable choice for a wide range of software development projects. Clean Architecture also allows for easy modification or replacement of any part of the program without affecting other parts of the program.

## Contact
Author: William Winkler <br>
LinkedIn: https://www.linkedin.com/in/williambwinkler/ <br>
Email: williambwinkler@gmail.com
