const brandSelect = document.getElementById('brandSelect');
const modelSelect = document.getElementById('modelSelect');

// Загрузка брендов
function loadBrands() {
    fetch('/get-brands')
        .then(res => res.json())
        .then(brands => {
            brandSelect.innerHTML = '<option value="">Выберите бренд</option>';
            brands.forEach(b => {
                const opt = document.createElement('option');
                opt.value = b.id;
                opt.textContent = b.name;
                brandSelect.appendChild(opt);
            });
        });
}

// Загрузка моделей
function loadModels() {
    const brandId = brandSelect.value;
    if (!brandId) {
        modelSelect.innerHTML = '<option value="">Выберите модель</option>';
        return;
    }
    fetch(`/get-models?brandId=${brandId}`)
        .then(res => res.json())
        .then(models => {
            modelSelect.innerHTML = '<option value="">Выберите модель</option>';
            models.forEach(m => {
                const opt = document.createElement('option');
                opt.value = m.id;
                opt.textContent = m.name;
                modelSelect.appendChild(opt);
            });
        });
}

// Инициализация
loadBrands();

// Отправка формы
document.getElementById('repairForm').addEventListener('submit', async e => {
    e.preventDefault();
    const formData = new FormData(e.target);

    const res = await fetch('/save-data', {
        method: 'POST',
        body: formData
    });

    if (res.ok) {
        alert('Автомобиль успешно добавлен!');
        e.target.reset();
    } else {
        const text = await res.text();
        alert('Ошибка: ' + text);
    }
});
