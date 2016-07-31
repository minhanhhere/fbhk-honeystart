F.onLocate = function(req) {
    if (req.query.lang === 'en')
        return 'en';
    return 'vi';
};