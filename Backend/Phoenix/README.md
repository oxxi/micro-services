# ElixirCactus

To start your Phoenix server:

- Run `mix setup` to install and setup dependencies
- Start Phoenix endpoint with `mix phx.server` or inside IEx with `iex -S mix phx.server`

Connect to WebSocket

- ws://localhost:4000/ws/websocket?vsn=2.0.0

client example

```JS
    webSocket.current = new WebSocket(url);

    webSocket.current.onopen = () => {
      webSocket.current?.send(JSON.stringify(['1', '1', 'broadcast:lobby', 'phx_join', {}]));
    };

    webSocket.current.onmessage = (event: MessageEvent<string>) => {
      // logic
    }
```
