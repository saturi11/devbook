
$('#form-cadastro').on('submit', criarUsuario);



function criarUsuario(evento) {
    evento.preventDefault();
    console.log("ta funcionando")

    if ($("#senha").val() != $("#confirmar-senha").val()) {
        Swal.fire('Oops...', 'as senhas nao sao iguais!', 'error')
        return;
    }
    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome: $("#nome").val(),
            email: $("#email").val(),
            nick: $("#nick").val(),
            senha: $("#senha").val(),
        }
    }).done(function () {
        Swal.fire('Sucesso!', 'Usuario cadastrado!', 'success').then(function () {
            $.ajax({
                url: "/login",
                method: "POST",
                data: {
                    email: $("#email").val(),
                    senha: $("#senha").val(),
                }
            }).done(function(){
                window.location = "/home"
            }).fail(function(){
                Swal.fire('Oops...', 'Erro ao autenticar usuario!', 'error')
            })
        })
    }).fail(function () {
        Swal.fire('Oops...', 'Erro ao cadastrar usuario!', 'error')
    });

}

