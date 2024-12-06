document.addEventListener("DOMContentLoaded", function() {
    document.querySelectorAll('p').forEach(function(paragraph) {
        let htmlContent = paragraph.innerHTML;

        htmlContent = htmlContent.replace(/`([^`]+?)`/g, '<span class="inline-code">$1</span>');
        
        if (htmlContent.startsWith('~&gt;')) {
            htmlContent = htmlContent.replace('~&gt;', '').trim();
            paragraph.classList.add('note-block', 'note-block-yellow');

            htmlContent = '<i class="fas fa-exclamation-triangle" style="color: #ffc107; margin-right: 0.5em;"></i>' + htmlContent;
        
        } else if (htmlContent.startsWith('-&gt;')) {
            htmlContent = htmlContent.replace('-&gt;', '').trim();
            paragraph.classList.add('note-block', 'note-block-blue');

            htmlContent = '<i class="fas fa-info-circle" style="color: #007acc; margin-right: 0.5em;"></i>' + htmlContent;

        } else if (htmlContent.startsWith('!&gt;')) {
            htmlContent = htmlContent.replace('!&gt;', '').trim();
            paragraph.classList.add('note-block', 'note-block-red');

            htmlContent = '<i class="fas fa-exclamation-circle" style="color: #d9534f; margin-right: 0.5em;"></i>' + htmlContent;
        }

        

        paragraph.innerHTML = htmlContent;
    });
});
