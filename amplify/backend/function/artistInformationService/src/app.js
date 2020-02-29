/*
Copyright 2017 - 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
Licensed under the Apache License, Version 2.0 (the "License"). You may not use this file except in compliance with the License. A copy of the License is located at
    http://aws.amazon.com/apache2.0/
or in the "license" file accompanying this file. This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and limitations under the License.
*/


var express = require('express');
var bodyParser = require('body-parser');
var awsServerlessExpressMiddleware = require('aws-serverless-express/middleware');
var fetch = require('node-fetch');

// declare a new express app
var app = express();
app.use(bodyParser.json());
app.use(awsServerlessExpressMiddleware.eventContext());

// Enable CORS for all methods
app.use(function (req, res, next) {
  res.header("Access-Control-Allow-Origin", "*");
  res.header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept");
  next()
});


/**********************
 * Example get method *
 **********************/

app.get('/api/v1/information/artists/:artist', async function (req, res) {
  const tokenResponse = await fetch('https://accounts.spotify.com/api/token',
    {
      method: 'POST',
      headers: {
        'Authorization': `Basic ${Buffer.from(process.env.SPOTIFY_CLIENT_ID + ':' + process.env.SPOTIFY_CLIENT_SECRET).toString('base64')}`,
        'Content-Type': 'application/x-www-form-urlencoded'
      },
      body: 'grant_type=client_credentials'
    });
  const tokenJson = await tokenResponse.json();
  const searchResponse = await fetch(`https://api.spotify.com/v1/search?q=${req.params.artist}&type=artist&limit=1`,
    {
      headers: { 'Authorization': `Bearer ${tokenJson.access_token}` }
    });
  const searchJson = await searchResponse.json();
  res.json({ image: searchJson.artists.items[0].images[1].url });
});

const server = app.listen(3000, function () {
});

app.stopServer = function () {
  server.close();
};

// Export the app object. When executing the application local this does nothing. However,
// to port it to AWS Lambda we will create a wrapper around that will load the app from
// this file
module.exports = app;
