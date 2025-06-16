function openModal(modalId) {
        const modal = document.getElementById(modalId);
        const form = modal.querySelector('form');

        // Animation d'apparition
        modal.classList.remove('hidden');
        setTimeout(() => {
            modal.classList.add('opacity-100');
            form.classList.add('scale-100');
        }, 10);

        function closeModal() {
            // Animation de fermeture
            modal.classList.remove('opacity-100');
            form.classList.remove('scale-100');

            // Attendre la fin de l'animation avant de masquer
            setTimeout(() => {
            modal.classList.add('hidden');
            }, 300);

            modal.removeEventListener('click', handleOutsideClick);
            document.removeEventListener('keydown', handleEsc);
        }

        function handleOutsideClick(e) {
            if (e.target === modal) closeModal();
        }

        function handleEsc(e) {
            if (e.key === 'Escape') closeModal();
        }

        // Attacher les événements
        modal.addEventListener('click', handleOutsideClick);
        document.addEventListener('keydown', handleEsc);

        // Bouton "Annuler"
        const cancelBtn = modal.querySelector('button[type="reset"]');
        if (cancelBtn) {
            cancelBtn.addEventListener('click', (e) => {
            e.preventDefault();
            closeModal();
            });
        }

        // Focus auto
        const firstInput = modal.querySelector('input, select, textarea');
        if (firstInput) firstInput.focus();
    }
