# Commands

- `/register <name><password>` - registration of a new user (Name and password will be stored in an accounts database).
- `/login <name><password>` - login of a user (name and password will be checked in an accounts database).
- `/list_rooms` - show the list of all active rooms.
- `/join <roomName>` - join the the room (or create a new room) to chat with users. If roomName is omitted the user will join the common room.
- `/write_to <userName>` - create a new private room with userName.
- `/list` - show the list of all connected users.
- `/list_accounts` - show the list of all registered users.
- `/change_name <name><password><newName>` - name and password will be checked and if it's OK then the newName will be accepted.
- `/msg <message>` - broadcast message to all connected users.
- `/quit_room` - finish the room and go back to the common room.
- `/quit` - disconnect from the chat server.
