module.exports = {
  async redirects() {
    return [
      {
        source: '/',
        destination: '/login',
        permanent: true,
      },
    ]
  },
  async rewrites(){
    return [
    {
      source: "/api",
      destination: "http://localhost:5006/api/",
    },
  ];
  }
}
