/**
 *  �ַ���ת�����ߺ���  (�շ�ʽ���»�������ʽ��ת)
 *
 * */

var fromCamel = function(objStr) {

    // the first word should be lowercase
    let str = objStr.replace(/^\S/g, (s) => {
        return s.toLowerCase();
    });

    // transform to underline style
    return str.replace(/([A-Z])/g, "_$1").toLowerCase();
};

var toCamel = function(objStr) {

    return objStr.replace(/_(\w)/g, (s) => {
        return s.slice(1).toUpperCase();
    });
};

module.exports = {
    fromCamel: fromCamel,
    toCamel: toCamel
};
