FROM node:12.6.0 AS build
WORKDIR /app

COPY . /app/

RUN yarn install \
      --frozen-lockfile

RUN yarn run build

# COMMENT TEST

FROM node:12.6.0-slim AS release
WORKDIR /app
COPY --from=build /app /app

ENTRYPOINT [ "yarn", "start" ]
