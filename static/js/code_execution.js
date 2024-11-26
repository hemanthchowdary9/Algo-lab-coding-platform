// Function to display test case results
function DisplayTestCaseResults(testCases) {

    document.getElementById("loading-spinner").style.display = "none";

    const resultsContainer = document.getElementById('test-case-results');
    const collapseButton = document.getElementById('collapse-button-id');

    // Clear previous results
    resultsContainer.innerHTML = '';

    console.log("testcases", testCases)
    // Check if there are test cases
    if (testCases.length > 0) {
        // Show the collapse button
        collapseButton.style.display = 'block';

        // Iterate over the test cases and create details elements
        testCases.forEach((result, index) => {
            const resultCard = document.createElement('div');
            resultCard.classList.add('result-card', result.isTestCasePassed ? 'success' : 'error');

            resultCard.innerHTML = `
        <details>
          <summary>Test Case ${index + 1}: ${result.isTestCasePassed ? 'Success' : 'Failed'}</summary>
          <div><strong>Input:</strong><pre>${result['input']}</pre></div>
          <div><strong>Expected Output:</strong><pre>${result['expectedOutput']}</pre></div>
          <div><strong>Your Output:</strong><pre>${result['output']}</pre></div>
          <div><strong>CPU Time:</strong> ${result.cpuTime}s</div>
          <div class="status ${result.isExecutionSuccess ? 'success' : 'error'}" style="display: none">
            ${result.isExecutionSuccess ? 'Execution Successful' : 'Execution Failed'}
          </div>
        </details>
      `;
            // Add event listener to manage collapsing behavior
            const detailsElement = resultCard.querySelector('details');
            detailsElement.addEventListener('toggle', function() {
                if (detailsElement.open) {
                    // Close all other details
                    document.querySelectorAll('details').forEach(otherDetails => {
                        if (otherDetails !== detailsElement) {
                            otherDetails.removeAttribute('open');
                        }
                    });
                }
            });

            resultsContainer.appendChild(resultCard);
        });
    } else {
        // Hide the collapse button if there are no test cases
        collapseButton.style.display = 'none';
    }
}

// Collapse all details when the button is clicked
document.getElementById('collapse-button-id').addEventListener('click', function() {
    const resultsContainer = document.getElementById('test-case-results');
    const collapseButton = document.getElementById('collapse-button-id');

    // Clear previous results
    resultsContainer.innerHTML = '';
    collapseButton.style.display = 'none';
});
