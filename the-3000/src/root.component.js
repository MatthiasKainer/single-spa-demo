import React from 'react'

export default function Root(props) {
  return (
    <section>
      <h1>Hello From React!</h1>
      {props.name} is mounted!
    </section>
  )
}