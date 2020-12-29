Minimalistic prometheus exporter for zendesk tickets count.

## Usage

~~~ shell
docker build -t micro_prometheus_zendesk_exporter .
docker run -e ZENDESK_DOMAIN=ZENDESK_DOMAIN -e ZENDESK_USER=ZENDESK_USER -e ZENDESK_PASSWORD=ZENDESK_PASSWORD -p "9803:9803" micro_prometheus_zendesk_exporter
~~~

Where `ZENDESK_DOMAIN` is part of your zendesk url (https://ZENDESK_DOMAIN.zendesk.com).

## Alternatives

There is a couple of similar projects: https://github.com/shakapark/Zendesk-Exporter and https://github.com/Tagnard/zendesk_exporter, but they load all tickets from zendesk (and it takes a while) when my implementation request only a count of tickets.