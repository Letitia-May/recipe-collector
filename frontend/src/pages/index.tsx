import Head from 'next/head'
import styled from 'styled-components'

const StyledHeader = styled.h1`
  color: grey;
`;

export default function Home() {
  return (
    <>
      <Head>
        <title>Recipe collector</title>
        <meta name="description" content="Collection of recipes that I love" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main>
        <StyledHeader>Recipe collector</StyledHeader>
      </main>
    </>
  )
}
