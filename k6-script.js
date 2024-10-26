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
    { duration: '5m', target: 1000 },
    { duration: '3m', target: 3000 },
    { duration: '5m', target: 7000 },
    { duration: '9m', target: 500 },
    { duration: '3m', target: 20 },
    { duration: '1m', target: 0 },
  ],
};

export default function () {
  // Iterate through each endpoint
  endpoints.forEach((url) => {
    const res = http.get(url);
  });
  sleep(1);
}
