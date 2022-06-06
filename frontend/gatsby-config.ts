import type { GatsbyConfig } from "gatsby";

const config: GatsbyConfig = {
  siteMetadata: {
    title: `recipe-collector`,
    siteUrl: `https://www.yourdomain.tld` //TODO: update
  },
  plugins: ["gatsby-plugin-styled-components"]
};

export default config;
