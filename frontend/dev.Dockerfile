FROM node:18.15.0
WORKDIR /app
COPY package*.json ./
COPY tsconfig.json ./
RUN npm install
EXPOSE 3000
ENV WATCHPACK_POLLING=true
ENTRYPOINT ["npm", "start"]
