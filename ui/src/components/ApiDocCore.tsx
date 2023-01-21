
import React, { useEffect, useRef } from 'react'
import RapiDocReact from './RapiDocReact';
import App from '../App';

type Props = {}

const ApiDocCore = (props: Props) => {

  const inputElement = useRef<any>();



  return (
    <RapiDocReact
      ref={inputElement}
      specLoaded={(spec) => { console.log(spec); }}
      show-header={true}
      allow-spec-file-load={false}
      allow-spec-url-load={false}
      bg-color="#ffff"
      primary-color="#386641"
      text-color={"#081c15"}
      header-color={'#081c15'}
      nav-bg-color={"#081c15"}
      spec-url="/spec-url"
      render-style="focused"

      theme="dark"
      style={{ height: '100vh', width: '100%' }}
    >

      <div slot='header'>
        <h3>Dok + cli</h3>
      </div>
    </RapiDocReact>
  )
}

export default ApiDocCore