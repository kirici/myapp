import http from 'k6/http';
import { check, sleep } from 'k6';

// Define the endpoints directly in the script
const endpoints = [
  `http://web:${__ENV.PORT}/`,
  `http://web:${__ENV.PORT}/work`,
  `http://web:${__ENV.PORT}/math`,
  `http://web:${__ENV.PORT}/isnotreal`
];

// Random rate generator simulating sine-like, out-of-phase waves
function getRandomizedInterval() {
  return Math.random() * 2 + 0.5; // Random interval between 0.5 and 2.5 seconds
}

// Define varying stages
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
    
    // Check for successful response
    check(res, {
      'status is 200': (r) => r.status === 200,
    });
  });

  // Pause between requests based on randomized interval
  sleep(getRandomizedInterval());
}
