// import http from 'k6/http';
// import { sleep } from 'k6';

// export let options = {
//     vus: 1,  // Number of virtual users
//     duration: '10s', // Duration of the test
// };

// const numberOfIterations = 5; // Set the desired number of iterations

// export default function () {
//     for (let i = 0; i < numberOfIterations; i++) {
//         // Your test logic here

//         // Example: Make an HTTP request
//         http.get('http://localhost:8090/accounts/0009933587');

//         // Sleep for a short duration (e.g., 1 second)
//         sleep(1);
//     }
// }

import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
    stages: [
        { duration: '3s', target: 20000 },    // 5 seconds with 10 VUs
        { duration: '3s', target: 7000 },   // 5 seconds with 50 VUs
        { duration: '3s', target: 12000 },   // 5 seconds with 100 VUs
        { duration: '3s', target: 6000 },  // 5 seconds with 50 VUs
        { duration: '3s', target: 9000 },  // 5 seconds with 10 VUs
    ],
};

export default function () {
    // http.get('http://127.0.0.1:8001/api/v1/namespaces/test-kube-namespace/services/test-kube-service/proxy/get-all-transaction');
    let response = http.get('http://192.168.73.2:30000/test');

    check(response, {
        'status is 200': (r) => r.status === 200,
    });

    sleep(1);
}

// k6 run load_test.js
