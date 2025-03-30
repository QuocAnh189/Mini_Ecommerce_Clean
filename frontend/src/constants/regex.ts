export const PASSWORD_REGEX = {
  upperCase: new RegExp('.*[A-Z].*'),
  lowerCase: new RegExp('.*[a-z].*'),
  number: new RegExp('.*\\d.*'),
  specialCharacter: new RegExp(
    '.*[`~<>?,./!@#$%^&*()\\-_+="\'|{}\\[\\];:\\\\].*'
  )
};

export const PHONE_REGEX = new RegExp(
  '^[+]?[(]?[0-9]{3}[)]?[-s.]?[0-9]{3}[-s.]?[0-9]{4,6}$'
);
