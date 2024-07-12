FROM node:20.15.0-alpine AS build

WORKDIR /usr/src/app

COPY yarn.lock package.json ./

RUN \
  --mount=id=yarn-cache,type=cache,target=/root/.yarn \
  YARN_CACHE_FOLDER=/root/.yarn \
  yarn install --frozen-lockfile --non-interactive --production=false

FROM gcr.io/distroless/nodejs20-debian12

WORKDIR /usr/src/app

COPY --from=build /usr/src/app/node_modules ./node_modules

CMD ["node_modules/.bin/next", "dev", "-p", "3100"]