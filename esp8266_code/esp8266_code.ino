#include <ESP8266WiFi.h>
#include "credentials.h"

const int switch1 = 14; // D5

WiFiServer server(80);
void setup() {
  Serial.begin(9600);
  Serial.print("Connecting to ");
  Serial.println(ssid);
  WiFi.begin(ssid, password);

  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }
  Serial.println("");
  Serial.println("WiFi connected");
  server.begin();
  Serial.print(WiFi.localIP());

  pinMode(switch1, OUTPUT);
}
void loop() {
  WiFiClient client = server.available();
  if (!client) {
    return;
  }
  while (!client.available()) {
    delay(1);
  }
  String request = client.readStringUntil('\r');
  client.flush();

  client.println("HTTP/1.1 200 OK");
  client.println("Content-Type: text/html");
  client.println("");

  if (request.indexOf("/open") != -1) {
    digitalWrite(switch1, HIGH);
    delay(100);
    digitalWrite(switch1, LOW);
    client.println("OK");
  } else {
    client.println("It's working!");
  }

  delay(1);
}