FROM node:22-alpine3.22

WORKDIR /app
COPY app /app
RUN npm install
EXPOSE 8002
CMD ["node", "index.js"]