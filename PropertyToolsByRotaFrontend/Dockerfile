FROM node:20-alpine

WORKDIR /app
COPY .out ./.out
COPY public ./public
COPY package.json ./
COPY node_modules ./node_modules

CMD ["pnpm", "start"]