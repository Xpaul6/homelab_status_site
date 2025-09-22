FROM node:21-alpine AS builder

WORKDIR /app

COPY package*.json ./
RUN npm ci

COPY . .

RUN npm run build

FROM node:21-alpine

WORKDIR /app

COPY package*.json ./
RUN npm ci

COPY --from=builder /app/dist ./dist
COPY vite.config.js .
COPY svelte.config.js .
COPY jsconfig.json .

EXPOSE 5555

CMD ["npm", "run", "preview"]
