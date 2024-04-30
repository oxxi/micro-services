defmodule ElixirCactusWeb.MessageChannel do
  use ElixirCactusWeb, :channel

  @impl true
  def join("message:lobby", payload, socket) do

    IO.inspect(payload, label: "en lobby")
    broadcast!(socket, "message", %{body: "hola"})
    {:noreply, socket}
  end

  @impl true
  def join("broadcast:lobby", _message, socket) do
    send(self(), :after_join)
    schedule_broadcast()
    {:ok, socket}
  end


  @impl true
  def handle_info(:after_join, socket) do
    fetch_and_broadcast(socket)
  end

  @impl true
  def handle_info(:broadcast_tick, socket) do
    fetch_and_broadcast(socket)
  end

  defp fetch_and_broadcast(socket) do
    case ElixirCactusWeb.Api.MetricsFetch.fetch do
      {:ok, info} ->
        broadcast(socket, "new", %{message: info})
      :error ->
        IO.puts("Error al intentar acceder a los datos")
    end
    {:noreply, socket}
  end

  defp schedule_broadcast() do
    :timer.send_interval(5_000, :broadcast_tick)  # send :broadcast_tick every 20 segundos
  end


  # Channels can be used in a request/response fashion
  # by sending replies to requests from the client
  @impl true
  def handle_in("ping", payload, socket) do
    {:reply, {:ok, payload}, socket}
  end

  # It is also common to receive messages from the client and
  # broadcast to everyone in the current topic (message:lobby).
  @impl true
  def handle_in("shout", payload, socket) do
    broadcast(socket, "shout", payload)
    {:noreply, socket}
  end

end
