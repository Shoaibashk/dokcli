import React from "react";
import "rapidoc";
import RapiDocReact from "./RapiDocReact";

export default function App() {
  return (
    <RapiDocReact
    specLoaded={(spec) => {
      console.log(spec);
    }}
    show-header={false}
    spec-url="https://petstore.swagger.io/v2/swagger.json"
    render-style="read"
    theme="dark"
    style={{ height: '100vh', width: '100%' }}
  />
  );
}