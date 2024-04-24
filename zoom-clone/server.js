const express = require('express');
const app = express();

const http = require('http');
const server = http.createServer(app);
const createIO = require('socket.io');

const io = createIO(server);
const { v4: uuidV4 } = require('uuid');

app.set('view engine', 'ejs');
app.use(express.static('public'));

app.get('/', (req, res) => {
  const dynRoomId = uuidV4();
  res.redirect(`/${dynRoomId}`);
});

app.get('/:roomId', (req, res) => {
  res.render('room', { roomId: req.params['roomId'] });
});

io.on(
  'connection',
  /** @type {(socket: import('socket.io').Socket) => void} */
  (socket) => {
    socket.on('join-room', (roomId, userId) => {
      console.log('join-room', roomId, userId);
      socket.join(roomId);
      socket.broadcast.to(roomId).emit('user-connected', userId);

      socket.on('disconnect', () => {
        socket.broadcast.to(userId).emit('user-disconnected', userId);
      });
    });
  }
);

const PORT = 3030;
server.listen(PORT, function () {
  console.log(`listening at ${PORT}`);
});
