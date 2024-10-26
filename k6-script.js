import http from 'k6/http';
import { check, sleep } from 'k6';

// Routes
const endpoints = [
  `http://web:${__ENV.PORT}/`,
  `http://web:${__ENV.PORT}/work`,
  `http://web:${__ENV.PORT}/math`,
  `http://web:${__ENV.PORT}/isnotreal`
];

// Stages
export const options = {
  stages: [
    { duration: '2m', target: 1000 },
    { duration: '1m', target: 2000 },
    { duration: '1m', target: 4000 },
    { duration: '1m', target: 8000 },
    { duration: '1m', target: 8000 },
    { duration: '0m', target: 500 },
    { duration: '1m', target: 500 },
    { duration: '3m', target: 5000 },
    { duration: '2m', target: 1000 },
  ],
};

export default function () {
  // Iterate through each endpoint
  endpoints.forEach((url) => {
    const res = http.get(url);
  });
  sleep(1);
}
