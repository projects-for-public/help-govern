document.addEventListener('DOMContentLoaded', function () {
    const form = document.getElementById('report-form');
    const resultDiv = document.getElementById('report-result');
    const categorySelect = document.getElementById('category');
    const latInput = document.getElementById('latitude');
    const lngInput = document.getElementById('longitude');

    // Placeholder: Populate categories (should fetch from backend in future)
    const categories = [
        'potholes',
        'broken_streetlight',
        'no_streetlight',
        'water_leaks',
        'poor_drainage',
        'damaged_sidewalk',
        'accident_prone',
        'garbage_heap',
        'wrong_side_driving'
    ];
    categories.forEach(cat => {
        const opt = document.createElement('option');
        opt.value = cat;
        opt.textContent = cat.replace(/_/g, ' ').replace(/\b\w/g, l => l.toUpperCase());
        categorySelect.appendChild(opt);
    });

    // Helper to reset error highlights
    function resetFieldStyles() {
        [categorySelect, latInput, lngInput].forEach(f => f.classList.remove('input-error'));
    }

    form.addEventListener('submit', async function (e) {
        e.preventDefault();
        resultDiv.textContent = '';
        resetFieldStyles();
        let hasError = false;
        let errorMsg = '';
        const data = {
            category: form.category.value,
            latitude: parseFloat(form.latitude.value),
            longitude: parseFloat(form.longitude.value),
            description: form.description.value
        };
        if (!data.category) {
            errorMsg += 'Please select a category.\n';
            categorySelect.classList.add('input-error');
            hasError = true;
        }
        if (isNaN(data.latitude) || data.latitude < -90 || data.latitude > 90) {
            errorMsg += 'Please enter a valid latitude (-90 to 90).\n';
            latInput.classList.add('input-error');
            hasError = true;
        }
        if (isNaN(data.longitude) || data.longitude < -180 || data.longitude > 180) {
            errorMsg += 'Please enter a valid longitude (-180 to 180).\n';
            lngInput.classList.add('input-error');
            hasError = true;
        }
        if (hasError) {
            resultDiv.innerHTML = `<span style='color:red;white-space:pre-line;'>${errorMsg}</span>`;
            return;
        }
        try {
            const resp = await fetch('/reports', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(data)
            });
            const respData = await resp.json();
            if (resp.ok) {
                resultDiv.innerHTML = `<span style='color:green'>${respData.message}</span><br>Share URL: <a href='${respData.share_url}' target='_blank'>${respData.share_url}</a>`;
                form.reset();
                resetFieldStyles();
            } else {
                let details = respData.details ? `<br><small>${respData.details}</small>` : '';
                resultDiv.innerHTML = `<span style='color:red'>${respData.error || 'Submission failed'}</span>${details}`;
            }
        } catch (err) {
            resultDiv.innerHTML = `<span style='color:red'>Network error: ${err}</span>`;
        }
    });

    // Remove error highlight on input
    [categorySelect, latInput, lngInput].forEach(f => {
        f.addEventListener('input', () => f.classList.remove('input-error'));
    });
}); 