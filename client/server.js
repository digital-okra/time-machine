const path = require("path");
const express = require("express");
const app = require("./public/App.js");

const server = express();

server.use(express.static(path.join(__dirname, "public")));

server.get("*", function(req, res) {
  const { html } = app.render({ url: req.url });

  res.write(`
  <!DOCTYPE html>
  <html>
  <head>
    <title>Time Machine</title>
    <meta charset="utf-8">

    <!-- Fonts and CSS -->
    <link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons">
    <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto:300,400,500,600,700">
    <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto+Mono">
    
    <link rel='stylesheet' href='/global.css'>
    <link rel='stylesheet' href='/bundle.css'>

    <meta name="viewport" content="width=device-width, initial-scale=1">

    <style>
      html, body {
        margin: 0;
        padding: 0;
      }
    </style>
  </head>
  <body>
    <app>${html}</app>
    <script type="text/javascript" src="/bundle.js"></script>
  </body>
  </html>
  `
  );

  res.end();
});

const port = 3000;
server.listen(port, () => console.log(`Listening on port ${port}`));
