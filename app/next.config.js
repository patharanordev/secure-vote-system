/** @type {import('next').NextConfig} */
const nextConfig = {
  trailingSlash: false,
  reactStrictMode: true, // Recommended for the `pages` directory, default in `app`.
  experimental: {
    appDir: true,
  },
  async rewrites() {
    return [
      {
        source: '/',
        destination: '/dashboard'
      },
      {
        source: '/register',
        destination: '/signup'
      },
      {
        source: '/signin',
        destination: '/login'
      },
    ]
  },
}

module.exports = nextConfig
