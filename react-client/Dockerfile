FROM node:12-alpine AS builder

COPY . ./react-client
WORKDIR /react-client
RUN npm install
CMD ["npm", "start"]