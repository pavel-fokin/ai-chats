# Monorepo

## Context

This project consists of several components, which are co-developed.
As a sole developer, I need a simple setup for this project
that helps to oversee the whole solution.

## Decision
I decided to use a monorepo approach to manage the project's source code.
The monorepo will be following the next structure.

```sh
/likeit     # Project name.
    /apps   # All components of the project.
    /docs   # Project documentation.
```

## Consequences

### Pros:

- **Holistic view:** Keeps all the code and information under one repo.
That makes it simpler to oversee the whole project.

### Cons:

- **Tooling:** It's harder to setup CI/CD pipelines and configure building of images.
- **Layout:** It can be harder to manage and scaleproject layout
with a time.