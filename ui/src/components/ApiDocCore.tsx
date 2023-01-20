
import React, { useEffect, useRef } from 'react'
import RapiDocReact from './RapiDocReact';

type Props = {}

const ApiDocCore = (props: Props) => {

  const inputElement = useRef<any>();



  return (
    <RapiDocReact
      ref={inputElement}
      specLoaded={(spec) => { console.log(spec); }}
      show-header={false}
      spec-url="/spec-url"
      render-style="read"
      theme="dark"
      style={{ height: '100vh', width: '100%' }}
    />
  )
}

export default ApiDocCore