FROM elixir:alpine

WORKDIR /app
COPY . .


RUN mix do deps.get, deps.compile


# App Port
EXPOSE 4000

# Default Command
CMD ["mix", "phx.server"]