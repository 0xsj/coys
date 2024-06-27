FROM node:22

WORKDIR /app/auth-service

COPY package*.json ./

RUN npm install

COPY . .

# EXPOSE 3333

CMD ["npm", "run", "dev"]