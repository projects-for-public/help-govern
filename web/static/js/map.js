document.addEventListener('DOMContentLoaded', function () {
    if (!window.L) return;
    // Center on India
    var map = L.map('map').setView([22.5937, 78.9629], 5);
    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '&copy; OpenStreetMap contributors'
    }).addTo(map);

    var marker = null;
    // On map click, set marker and update form fields
    map.on('click', function (e) {
        var lat = e.latlng.lat.toFixed(6);
        var lng = e.latlng.lng.toFixed(6);
        if (marker) {
            marker.setLatLng(e.latlng);
        } else {
            marker = L.marker(e.latlng, { draggable: true }).addTo(map);
            marker.on('dragend', function (ev) {
                var pos = ev.target.getLatLng();
                setLatLngFields(pos.lat, pos.lng);
            });
        }
        setLatLngFields(lat, lng);
    });

    // Helper to set form fields
    function setLatLngFields(lat, lng) {
        var latField = document.getElementById('latitude');
        var lngField = document.getElementById('longitude');
        if (latField && lngField) {
            latField.value = lat;
            lngField.value = lng;
        }
    }

    // If form fields are filled, show marker
    var latField = document.getElementById('latitude');
    var lngField = document.getElementById('longitude');
    if (latField && lngField && latField.value && lngField.value) {
        var lat = parseFloat(latField.value);
        var lng = parseFloat(lngField.value);
        if (!isNaN(lat) && !isNaN(lng)) {
            marker = L.marker([lat, lng], { draggable: true }).addTo(map);
            map.setView([lat, lng], 14);
            marker.on('dragend', function (ev) {
                var pos = ev.target.getLatLng();
                setLatLngFields(pos.lat, pos.lng);
            });
        }
    }

    // --- Display existing reports as markers with clustering ---
    // TODO: Replace this mock data with a real API call to /reports/geo or similar
    const mockReports = [
        {
            id: 1,
            category: 'potholes',
            latitude: 28.6139,
            longitude: 77.2090,
            description: 'Large pothole near Connaught Place',
            status: 'pending'
        },
        {
            id: 2,
            category: 'broken_streetlight',
            latitude: 19.0760,
            longitude: 72.8777,
            description: 'Streetlight not working on Marine Drive',
            status: 'in_progress'
        },
        {
            id: 3,
            category: 'water_leaks',
            latitude: 13.0827,
            longitude: 80.2707,
            description: 'Water leak near Anna Salai',
            status: 'resolved'
        }
    ];

    var markers = L.markerClusterGroup();
    mockReports.forEach(function (report) {
        var m = L.marker([report.latitude, report.longitude]);
        m.bindPopup(
            `<b>${report.category.replace(/_/g, ' ')}</b><br>` +
            `${report.description}<br>` +
            `<span>Status: ${report.status}</span>`
        );
        markers.addLayer(m);
    });
    map.addLayer(markers);

    // --- GPS Location Capture ---
    // Add a button to the map for geolocation
    var locateBtn = L.control({ position: 'topleft' });
    locateBtn.onAdd = function (map) {
        var div = L.DomUtil.create('div', 'leaflet-bar leaflet-control leaflet-control-custom');
        div.style.backgroundColor = '#fff';
        div.style.cursor = 'pointer';
        div.style.width = '34px';
        div.style.height = '34px';
        div.style.display = 'flex';
        div.style.alignItems = 'center';
        div.style.justifyContent = 'center';
        div.title = 'Detect my location';
        div.innerHTML = '<span style="font-size:20px;">📍</span>';
        div.onclick = function (e) {
            e.stopPropagation();
            if (!navigator.geolocation) {
                alert('Geolocation is not supported by your browser.');
                return;
            }
            div.innerHTML = '<span style="font-size:20px;">⏳</span>';
            navigator.geolocation.getCurrentPosition(function (pos) {
                var lat = pos.coords.latitude;
                var lng = pos.coords.longitude;
                map.setView([lat, lng], 16);
                if (marker) {
                    marker.setLatLng([lat, lng]);
                } else {
                    marker = L.marker([lat, lng], { draggable: true }).addTo(map);
                    marker.on('dragend', function (ev) {
                        var p = ev.target.getLatLng();
                        setLatLngFields(p.lat, p.lng);
                    });
                }
                setLatLngFields(lat, lng);
                div.innerHTML = '<span style="font-size:20px;">📍</span>';
            }, function (err) {
                alert('Could not get your location: ' + err.message);
                div.innerHTML = '<span style="font-size:20px;">📍</span>';
            });
        };
        return div;
    };
    locateBtn.addTo(map);
}); 