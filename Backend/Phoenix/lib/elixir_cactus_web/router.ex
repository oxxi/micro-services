defmodule ElixirCactusWeb.Router do
  use ElixirCactusWeb, :router

  pipeline :api do
    plug :accepts, ["json"]
  end

  scope "/api", ElixirCactusWeb do
    pipe_through :api
  end
end
