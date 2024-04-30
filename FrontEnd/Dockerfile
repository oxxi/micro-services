FROM node:18-alpine AS IMAGE_BUILDER
## create work directory
WORKDIR /app
## copy package.json
COPY package.json .
## install depencencies 
RUN npm install
## copy all files excluding files in dockerignore
COPY . .
## expose port
EXPOSE 3000
## run add
CMD [ "npm", "run", "dev" ]


