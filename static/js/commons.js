
function toggleDropdown() {
    document.getElementById("dropdownContent").classList.toggle("show");
}

window.onclick = function(event) {
    if (!event.target.matches('.user-icon img')) {
        var dropdowns = document.getElementsByClassName("dropdown-content");
        for (var i = 0; i < dropdowns.length; i++) {
            var openDropdown = dropdowns[i];
            if (openDropdown.classList.contains('show')) {
                openDropdown.classList.remove('show');
            }
        }
    }
}


function ShowExplore() {
    // Logic to show challenges can be added later
    window.location.href = '/dashboard'
}


function ShowChallenges() {
    window.location.href = '/challenges'
}


function ShowContact() {
    window.location.href = '/customer-care'
}


function PurchasePremium() {
    console.log("Inside purchase premium")
    window.location.href = '/create-checkout-session'
}


function ShowSubmissions() {
    window.location.href = '/submissions'
}