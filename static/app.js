async function fetchAssets() {
    try {
        const response = await fetch('/assets');
        const assets = await response.json();
        const container = document.getElementById('assets');
        if (!assets.length) {
            container.innerHTML = '<p>No assets found.</p>';
            return;
        }
        const table = document.createElement('table');
        const header = table.insertRow();
        header.innerHTML = '<th>ID</th><th>Item Name</th><th>Barcode</th>';
        assets.forEach(asset => {
            const row = table.insertRow();
            const idCell = row.insertCell();
            idCell.textContent = asset.id_asset;
            const nameCell = row.insertCell();
            nameCell.textContent = asset.item_name;
            const barcodeCell = row.insertCell();
            barcodeCell.textContent = asset.no_barcode;
        });
        container.appendChild(table);
    } catch (error) {
        console.error('Error fetching assets:', error);
        document.getElementById('assets').innerText = 'Failed to load assets.';
    }
}

window.addEventListener('DOMContentLoaded', fetchAssets); 