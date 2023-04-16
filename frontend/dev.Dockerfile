FROM node:18.15.0
WORKDIR /app
COPY package*.json ./
COPY tsconfig.json ./
RUN npm install
EXPOSE 8081
ENV PORT=8081
ENTRYPOINT ["npm", "start"]
