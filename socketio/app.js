const axios = require('axios');
const app = require('express')();
const server = require('http').Server(app);
const io = require('socket.io')(server);

server.listen(process.env.PORT, () => {
  console.log('Server listening at port %d', process.env.PORT);
});

app.get('/', (req, res) => {
  res.sendFile(__dirname + '/index.html');
});

app.get('/messages', async (req, res) => {
  const { data } = await axios.get(process.env.API_HOST);
  res.send(data);
});

io.on('connection', (socket) => {
  socket.on('username', (data) => {
    socket.username = data.name;
  });

  socket.on('message', (data) => {
    data.username = socket.username;
    socket.broadcast.emit('message', data);
    axios.post(process.env.API_HOST, data);
  });
});
