defmodule ElixirCactusWeb.Api.MetricsFetch do
  require Logger

  def fetch do
    api_url = System.get_env("API_URL")
    IO.inspect(api_url,label: "URL golang service")
    case HTTPoison.get(api_url) do
      {:ok, %HTTPoison.Response{status_code: 200, body: body}} ->
        {:ok, Jason.decode!(body)}
      {:ok, %HTTPoison.Response{status_code: status_code}} ->
        Logger.error("Failed to fetch metrics, status code: #{status_code}")
        :error
      {:error, %HTTPoison.Error{reason: reason}} ->
        Logger.error("HTTP error: #{inspect(reason)}")
        :error
    end
  end
end
