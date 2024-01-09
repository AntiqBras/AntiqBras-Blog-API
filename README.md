# Antiqbras-Blog API

## Overview

This simple API enables writers to manage blog posts and their accounts. It supports operations like creating, editing, and deleting posts, as well as managing writer accounts.

## Features

### Writers
- **Management**: Writers can create, edit, and delete blog posts.
- **Account Creation**: Writers have the ability to create other writer accounts.
- **Default Writer**: A default writer account is always present upon application setup. The details of this default writer can be modified in the ".env" file (refer to ".env.example").

### Posts
- **Structure**: Each post contains a title, subtitle, content, slug, image, and the associated writer.
- **Accessibility**: Posts are publicly viewable. However, creation, editing, or deletion of posts requires writer authentication.

## Routes

### Posts

- **GET "/posts"**: Fetches all posts. Supports query parameters: "page" (page number) and "pageSize" (number of posts per page, default is 5).

- **GET "/posts/:slug"**: Retrieves detailed information of a post using its slug.

- **POST "/posts"**: Creates a new post; requires authentication.
  ```json
  {
    "hero_image": "https://i.imgur.com/cool_image",
    "title": "My first Post",
    "subtitle": "Best subtitle",
    "content": "Lorem Ipsum"
  }
  ```

- **PUT "/posts/:id"**: Edits a specific post by its ID; requires authentication.

- **DELETE "/posts/:id"**: Deletes a specific post by its ID; requires authentication.

### Writer

- **POST "/writers/login"**: Generates a JWT token for authenticating certain routes. Requires valid writer credentials. A default writer can be set up in the ".env" file.
  ```json
  {
    "username": "example_writer",
    "password": "123"
  }
  ```

- **POST "/writers/register"**: Registers a new writer; requires authentication.
  ```json
  {
    "username": "new_writer",
    "author": "John Doe",
    "password": "123"
  }
  ```

- **GET "/writers/:username"**: Retrieves information about a specific writer using their username; requires authentication.

## Environment Setup

- **Docker**: The application uses Docker. Database configurations can be set in the ".env" file (see ".env.example").

### Running the Database

```bash
docker-compose -f ./docker-compose.yml up
```

### Running the API

```bash
go run .
```

## To-Do

- [ ] Implement roles for writers.
- [x] Add pagination